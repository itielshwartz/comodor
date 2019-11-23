import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import ChipList from './ChipList.js';
import SelectSearch from 'react-select-search';
import '../assets/css/select-search.css';

const styles = theme => ({
  bar: {
    paddingTop: 16,
    paddingBottom: 16,
    marginBottom: 16
  }
});

var filters = [
    {
        id: "status",
        name: 'Status',
        items: [
            {name: 'deployed', value: 'status_deployed', isAvailable: true},
            {name: 'pending', value: 'status_pending', isAvailable: true},
        ]
    },
    {
        id: "cluster",
        name: 'Cluster',
        items: [
            {name: 'main', value: 'cluster_main', isAvailable: true},
        ]
    },
    {
        id: "namespace",
        name: 'Namespace',
        items: [
            {name: 'prod', value: 'namespace_prod', isAvailable: true},
            {name: 'staging', value: 'namespace_staging', isAvailable: true},
        ]
    },
];

class FilterBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      allFilters: filters,
      activeFilters: [],
    };
  }

  setItemAvailability = (key, isAvailable) => {
    var allFilters = this.state.allFilters,
      i = 0,
      j = 0,
      didFind = false;

    while (!didFind && i < allFilters.length) {
      j = 0;
      while (!didFind && j < allFilters[i].items.length) {
        didFind = allFilters[i].items[j].value === key;
        
        if (didFind) {
          allFilters[i].items[j].isAvailable = isAvailable;
        } else {
          j++;
        }
      }
      i += !didFind ? 1 : 0;
    }

    return { section: i, item: j }
  }

  filterAdded = (selected) => {
    var itemIndexPath = this.setItemAvailability(selected.value, false);

    // Adding to filters
    var activeFilters = this.state.activeFilters;
    const filterGroup = this.state.allFilters[itemIndexPath.section];
    activeFilters.push({
      name: filterGroup.name + ': ' + filterGroup.items[itemIndexPath.item].name,
      value: selected.value,
      key: selected.value
    });

    this.setState(state => ({ activeFilters }));
    this.props.onChange(activeFilters);
  };

  getAvailableFilters = () => {
    var allFilters = this.state.allFilters,
      availableFilters = [];

    for (var i in allFilters) {
      var group = {
        type: "group",
        name: allFilters[i].name,
        items: []
      };
      
      for (var j in allFilters[i].items) {
        if (allFilters[i].items[j].isAvailable) {
          group.items.push(allFilters[i].items[j])
        }
      }

      availableFilters.push(group);
    }

    return availableFilters;
  }

  filterRemoved = (key) => {
    this.setItemAvailability(key, true)

    // Removing from filters
    var activeFilters = this.state.activeFilters,
      i = 0,
      didFind = false;

    while (!didFind && i < activeFilters.length) {
      if (activeFilters[i].key === key) {
        didFind = true;
        activeFilters.splice(i, 1);
      } else {
        i++;
      }
    }

    this.setState(state => ({ activeFilters }));
    this.props.onChange(activeFilters);
  }

  render() {
    const { classes } = this.props;

    var availableFilters = this.getAvailableFilters();

    return (
      <div className={"box-h box-ic " + classes.bar}>
        <SelectSearch
          options={availableFilters}
          value=""
          name="filter"
          placeholder="Filter"
          onChange={this.filterAdded}
        />
        <ChipList
          options={this.state.activeFilters}
          onChange={this.filterRemoved}
        />
      </div>
    );
  }
}

export default withStyles(styles, { withTheme: true })(FilterBar);