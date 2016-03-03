/* @flow */

import React, {Component} from 'react'
import {Icon} from '../common-adapters'
import {globalStyles, globalColors} from '../styles/style-guide'
import type {Props} from './header'
import Text from './text'

export default class Header extends Component {
  props: Props;

  render () {
    return (
      <div style={{...this.props.style, ...styles.container}}>
        {this.props.children}
        {this.props.icon && <Icon type='logo-24' />}
        <Text type='Body' style={{flex: 1}}>{this.props.title}</Text>
        {this.props.onClose && (
          <div style={styles.close} onClick={() => this.props.onClose()}>
            <i className='fa fa-times' ></i>
          </div>
        )}
      </div>
    )
  }
}

Header.propTypes = {
  icon: React.PropTypes.bool,
  children: React.PropTypes.any,
  title: React.PropTypes.string,
  onClose: React.PropTypes.func,
  style: React.PropTypes.object
}

const styles = {
  container: {
    ...globalStyles.flexBoxRow,
    ...globalStyles.windowDragging,
    ...globalStyles.noSelect,
    paddingLeft: 10,
    paddingRight: 10,
    alignItems: 'center',
    height: 35
  },
  logo: {
    width: 22,
    height: 22,
    marginRight: 8
  },
  close: {
    ...globalStyles.flexBoxRow,
    ...globalStyles.clickable,
    ...globalStyles.windowDraggingClickable,
    color: globalColors.grey4,
    fontSize: 16,
    alignItems: 'flex-start',
    justifyContent: 'flex-end',
    width: 30,
    height: 24
  }
}
