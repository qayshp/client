import * as Kb from '../../common-adapters'
import * as React from 'react'
import * as Styles from '../../styles'
import PeopleItem, {TaskButton} from '../item'
import {WotStatusType} from '../../constants/types/rpc-gen'

type Props = {
  userForIcon: string
  otherUser: string
  voucher: string
  key: string
  vouchee: string
  wotStatus: WotStatusType
  badged: boolean
  onClickUser: (username: string) => void
  onDismiss: (voucher: string, vouchee: string) => void
}

const makeButtons = (props: Props): Array<TaskButton> => {
  const dismissButton = {
    label: 'Dismiss',
    mode: 'Secondary',
    onClick: () => props.onDismiss(props.voucher, props.vouchee),
  } as TaskButton
  switch (props.wotStatus) {
    case WotStatusType.proposed:
      return [
        {
          label: 'Review claim',
          onClick: () => props.onClickUser(props.vouchee),
        },
        dismissButton,
      ]
    case WotStatusType.accepted:
      return [
        {
          label: 'Review claim',
          onClick: () => props.onClickUser(props.vouchee),
        },
        dismissButton,
      ]
    case WotStatusType.rejected:
      return [
        {
          label: 'Edit claim',
          onClick: () => props.onClickUser(props.vouchee),
        },
        dismissButton,
      ]
    default:
      return []
  }
}

const makeMessage = (props: Props) => {
  const connectedUsernamesProps = {
    colorFollowing: true,
    inline: true,
    joinerStyle: {
      fontWeight: 'normal',
    },
    onUsernameClicked: 'profile',
    type: 'BodyBold',
    underline: true,
  } as const
  const voucherComponent = (
    <Kb.ConnectedUsernames
      {...connectedUsernamesProps}
      usernames={props.voucher}
      onUsernameClicked={props.onClickUser}
    />
  )
  const voucheeComponent = (
    <Kb.ConnectedUsernames
      {...connectedUsernamesProps}
      usernames={props.vouchee}
      onUsernameClicked={props.onClickUser}
    />
  )
  switch (props.wotStatus) {
    case WotStatusType.proposed:
      return <Kb.Text type="Body">{voucherComponent} submitted an entry to your web of trust.</Kb.Text>
    case WotStatusType.accepted:
      return <Kb.Text type="Body">{voucheeComponent} accepted your entry into their web of trust.</Kb.Text>
    case WotStatusType.rejected:
      return <Kb.Text type="Body">{voucheeComponent} rejected your entry into their web of trust.</Kb.Text>
    default:
      return <Kb.Text type="Body">unknown.</Kb.Text>
  }
}

const WotTask = (props: Props) => {
  return (
    <Kb.ClickableBox onClick={() => props.onClickUser(props.vouchee)}>
      <PeopleItem
        badged={props.badged}
        icon={
          <Kb.Avatar
            username={props.otherUser}
            onClick={() => props.onClickUser(props.otherUser)}
            size={Styles.isMobile ? 48 : 32}
          />
        }
        buttons={makeButtons(props)}
      >
        {makeMessage(props)}
      </PeopleItem>
    </Kb.ClickableBox>
  )
}

export default WotTask
