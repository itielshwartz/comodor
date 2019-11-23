import React from 'react';
import Chip from '@material-ui/core/chip';
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
  chipList: {
    '& .MuiChip-root': {
      marginLeft: 8
    }
  }
});

class ChipList extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      items: props.options
    };
  }

  onDelete = chip => {
    this.props.onChange(chip)
  };

  render() {
    const { classes } = this.props;

    return (
      <div className={"box-h box-ic " + classes.chipList}>
        {this.state.items.map(item => {
          return (
            <Chip
              key={item.key}
              label={item.name}
              onDelete={this.onDelete.bind(this, item.key)}
            />
          );
        })}
      </div>
    );
  }
}

export default withStyles(styles, { withTheme: true })(ChipList);