import classnames from 'classnames';
import * as React from 'react';
import { connect } from 'react-redux';

import { Bubble, ButtonBar, Segment, SegmentProps } from '@tea/component';
import { FormPanel } from '@tencent/ff-component';
import { bindActionCreators } from '@tencent/ff-redux';
import { t, Trans } from '@tencent/tea-app/lib/i18n';
import { SegmentOption } from '@tencent/tea-component/lib/segment/SegmentOption';

import { allActions } from '../actions';
import { validatorActions } from '../actions/validatorActions';
import { logModeList, originModeList } from '../constants/Config';
import { ContainerLogs } from '../models';
import { EditOriginContainerItemPanel } from './EditOriginContainerItemPanel';
import { ListOriginContainerItemPanel } from './ListOriginContainerItemPanel';
import { RootProps } from './LogStashApp';

/** 日志源的相关提示 */
const originModeTip = {
  selectAll: t('选择所有Namespace下的所有容器'),
  selectOne: t('选择指定Namespace下的容器')
};

// 判断是否能够新增Namespace
export function isCanAddContainerLog(containerLogs: ContainerLogs[], namespaceCount: number) {
  let canAdd = true,
    canSave = true,
    tip = t('请先完成待编辑项'),
    editingContainerLog = containerLogs.find(item => item.status === 'editing');

  if (editingContainerLog) {
    canSave = validatorActions._canAddContainerLog(editingContainerLog, containerLogs);
  }

  if (containerLogs.length >= namespaceCount) {
    canAdd = false;
    tip = t('指定容器日志源数量不可超过当前集群下Namespace的数量');
  }

  if (editingContainerLog && canAdd) {
    canAdd = canAdd && canSave;
  }

  return { canAdd, tip, canSave };
}

const mapDispatchToProps = dispatch =>
  Object.assign({}, bindActionCreators({ actions: allActions }, dispatch), {
    dispatch
  });

@connect(state => state, mapDispatchToProps)
export class EditOriginContainerPanel extends React.Component<RootProps, any> {
  componentDidMount(): void {
    let { actions } = this.props;
    if (window.location.href.includes('/tkestack-project')) {
      // 业务侧不显示"所有容器"
      actions.editLogStash.selectAllNamespace('selectOne');
    }
  }

  render() {
    let { actions, logStashEdit, namespaceList } = this.props,
      { isSelectedAllNamespace, logMode, containerLogs } = logStashEdit;

    // 日志源的类型
    let selectedOriginMode = originModeList.find(item => item.value === isSelectedAllNamespace);

    // 判断是否能添加命名空间
    let { canAdd, tip } = isCanAddContainerLog(containerLogs, namespaceList.data.recordCount);

    const originModeListSegments: SegmentOption[] = originModeList.map(mode => {
      return { value: mode.value, text: mode.name };
    });
    return (
      <FormPanel.Item
        label={t('日志源')}
        isShow={logMode === logModeList.container.value}
        message={originModeTip[isSelectedAllNamespace]}
      >
        {window.location.href.includes('/tkestack-project') ||
          <Segment
            options={originModeListSegments}
            value={selectedOriginMode.value}
            onChange={value => actions.editLogStash.selectAllNamespace(value)}
          />}

        {isSelectedAllNamespace === 'selectOne' && this._renderContainerLogList()}

        {window.location.href.includes('/tkestack-project') ||
        isSelectedAllNamespace === 'selectOne' && (
          <Bubble content={!canAdd ? tip : null} placement="right">
            <a
              href="javascript:;"
              className={classnames('add-btn', { disabled: !canAdd })}
              onClick={() => {
                canAdd && actions.editLogStash.addContainerLog();
              }}
              style={{ marginTop: '10px' }}
            >
              {t('添加Namespace')}
            </a>
          </Bubble>
        )}
      </FormPanel.Item>
    );
  }

  /** 渲染指定容器的内容 */
  private _renderContainerLogList() {
    let { logStashEdit } = this.props;
    return logStashEdit.containerLogs.map((containerLog, index) => {
      return containerLog.status === 'edited' ? (
        <ListOriginContainerItemPanel cKey={containerLog.id + ''} key={index} />
      ) : (
        <EditOriginContainerItemPanel cKey={containerLog.id + ''} key={index} />
      );
    });
  }
}
