import React from 'react';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
  button: {
    color: 'red',
    padding: 16
  }
});

function clicked() {
  console.log("!")
}

class FilterButton extends React.Component {
  render() {
    const { classes } = this.props;

    return (
      <button className={classes.button} onClick={clicked}>Filter</button>
    );
  }
}

// <Chip label="he" />

export default withStyles(styles, { withTheme: true })(FilterButton);