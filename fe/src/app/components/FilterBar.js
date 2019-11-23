import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Chip from '@material-ui/core/chip';
import FilterButton from './FilterButton.js';
import SelectSearch from 'react-select-search';
import '../assets/css/select-search.css';

const styles = theme => ({
  bar: {
    padding: 16
  }
});

const options = [
    {
        type: 'group',
        name: 'Status',
        items: [
            {name: 'deployed', value: 'status_deployed'},
            {name: 'pending', value: 'status_pending'},
        ]
    },
    {
        type: 'group',
        name: 'Cluster',
        items: [
            {name: 'main', value: 'cluster_main'},
        ]
    },
    {
        type: 'group',
        name: 'namespace',
        items: [
            {name: 'prod', value: 'namespace_prod'},
            {name: 'staging', value: 'namespace_staging'},
        ]
    },
];

function filterAdded(selected) {
  let raw = selected.value.split("_");
  let field = raw[0];
  let value = raw[1];
}

class FilterBar extends React.Component {
  constructor(props) {
    super(props);
    this.activeFilters = [];
  }

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.bar}>
        <SelectSearch options={options} value="" name="filter" placeholder="Filter" onChange={filterAdded} />
      </div>
    );
  }
}

export default withStyles(styles, { withTheme: true })(FilterBar);