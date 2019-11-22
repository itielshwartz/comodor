import React from 'react';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
  statusIndicator: {
    color: '#8a8a8a',
    fontSize: 12,

    '&::before': {
      display: 'inline-block',
      content: '""',
      width: 8,
      height: 8,
      borderRadius: 100,
      marginRight: 4
    },

    '&.good::before': { background: '#64dd17' },
    '&.pending::before': { background: '#00b8d4' },
    '&.bad::before': { background: '#ff3d00' },
    '&.unknown::before': { background: '#b0bec5' },
    '&.good': { color: '#64dd17' },
    '&.pending': { color: '#00b8d4' },
    '&.bad': { color: '#ff3d00' },
    '&.unknown': { color: '#b0bec5' }
  }
});

function synthesizeColor(status) {
  switch(status) {
    case "deployed":
      return "good";
    case "unknown":
      return "unknown";
    case "deleted":
    case "superseded":
    case "failed":
    case "deleting":
      return "bad";
    case "pending_install":
    case "pending_upgrade":
    case "pending_rollback":
      return "pending";
    default:
      return "";
  }
}

function synthesizeName(status) {
  switch(status) {
    case "deployed":
      return "Live";
    case "unknown":
      return "Unknown";
    case "deleted":
    case "superseded":
    case "failed":
    case "deleting":
      return "Down";
    case "pending_install":
    case "pending_upgrade":
    case "pending_rollback":
      return "Pending";
    default:
      return "";
  }
}

class StatusIndicator extends React.Component {


  render() {
    const { classes } = this.props;

    var color = synthesizeColor(this.props.status);
    var name = synthesizeName(this.props.status);

    return (
      <div className={"box-h box-ic " + classes.statusIndicator + " " + color}>
        <span>{name}</span>
      </div>
    );
  }
}

export default withStyles(styles, { withTheme: true })(StatusIndicator);