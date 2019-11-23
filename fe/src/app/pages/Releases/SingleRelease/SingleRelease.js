import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import Box from '@material-ui/core/Box';

import { makeStyles } from '@material-ui/core/styles';
import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import Typography from '@material-ui/core/Typography';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import { Services } from './Services/Services';
import { Deployments } from './Deployments/Deployments';
import { StatefulSets } from './StatefulSets/StatefulSets';
import { Pods } from './Pods/Pods';
import { History } from './History/History';

const GET_RELEASE = id => gql`
  {
    comodor_releases(where: {row_id: {_eq: ${id}}}) {
      cluster
      created_at
      name
      namespace
      revision
      row_id
      status
    }
  }
`;

const GET_SERVICES = (name, namespace, revision) => gql`
  {
    comodor_services(where: {unique_release_name: {_eq: "${name}"}, unique_release_namespace: {_eq: "${namespace}"}, unique_release_revision: {_eq: ${revision}}}) {
      cluster_ip
      created_at
      external_ip
      name
      ports
      type
    }
  }
`;

const ReleaseDetails = ({ release }) => {
  return (
    <>
    <Box display="flex" alignItems="flex-start">
      <Box p={1}>
        Cluster:
      </Box>
      <Box p={1}>
        {release.cluster}
      </Box>
    </Box>
    <Box display="flex" alignItems="flex-start">
      <Box p={1}>
        Status:
      </Box>
      <Box p={1}>
        {release.status}
      </Box>
    </Box>
    <Box display="flex" alignItems="flex-start">
      <Box p={1}>
        Revision:
      </Box>
      <Box p={1}>
        {release.revision}
      </Box>
    </Box>
    <Box display="flex" alignItems="flex-start">
      <Box p={1}>
        Created At:
      </Box>
      <Box p={1}>
        {release.created_at}
      </Box>
    </Box>
    </>
  )
}

const useStyles = makeStyles(theme => ({
  root: {
    width: '100%',
  },
  heading: {
    fontSize: theme.typography.pxToRem(15),
    flexBasis: '33.33%',
    flexShrink: 0,
  },
  secondaryHeading: {
    fontSize: theme.typography.pxToRem(15),
    color: theme.palette.text.secondary,
  },
}));

export const ReleaseData = ( {release} ) =>  {
  const classes = useStyles();
  const [expanded, setExpanded] = React.useState(false);

  const handleChange = panel => (event, isExpanded) => {
    setExpanded(isExpanded ? panel : false);
  };

  return (
      <div style={{marginLeft: 300}} >
        <div style={{marginLeft: 400}}>{`Release: ${release.name} ${release.namespace}`}</div>
        <ReleaseDetails release={release} />
        <div className={classes.root}>
      <ExpansionPanel expanded={expanded === 'panel1'} onChange={handleChange('panel1')}>
        <ExpansionPanelSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel1bh-content"
          id="panel1bh-header"
        >
          <Typography className={classes.heading}>Services</Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <Services release={release} />
        </ExpansionPanelDetails>
      </ExpansionPanel>
      <ExpansionPanel expanded={expanded === 'panel2'} onChange={handleChange('panel2')}>
        <ExpansionPanelSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel2bh-content"
          id="panel2bh-header"
        >
          <Typography className={classes.heading}>Deployments</Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <Deployments release={release} />
        </ExpansionPanelDetails>
      </ExpansionPanel>
      <ExpansionPanel expanded={expanded === 'panel3'} onChange={handleChange('panel3')}>
        <ExpansionPanelSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel3bh-content"
          id="panel3bh-header"
        >
          <Typography className={classes.heading}>Pods</Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <Pods release={release} />
        </ExpansionPanelDetails>
      </ExpansionPanel>
      <ExpansionPanel expanded={expanded === 'panel4'} onChange={handleChange('panel4')}>
        <ExpansionPanelSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel3bh-content"
          id="panel3bh-header"
        >
          <Typography className={classes.heading}>StatefulSets</Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <StatefulSets release={release} />
        </ExpansionPanelDetails>
      </ExpansionPanel>
      <ExpansionPanel expanded={expanded === 'panel5'} onChange={handleChange('panel5')}>
        <ExpansionPanelSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel3bh-content"
          id="panel3bh-header"
        >
          <Typography className={classes.heading}>History</Typography>
        </ExpansionPanelSummary>
        <ExpansionPanelDetails>
          <History release={release} />
        </ExpansionPanelDetails>
      </ExpansionPanel>
    </div>
    </div>
  )
}

export const SingleRelease = ({  match }) =>  {
  const { loading, error, data } = useQuery(GET_RELEASE(match.params.id));

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  const release = data.comodor_releases[0];

  return (
    <ReleaseData release={release} />
  )
}