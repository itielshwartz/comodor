import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Button } from '@material-ui/core';


const useStyles = makeStyles({
    root: {
      width: '100%',
      overflowX: 'auto',
    },
    table: {
      minWidth: 650,
    },
  });

const GET_HISTORY = (name, namespace, cluster, currentRevision) => gql`
  {
    comodor_app_history(order_by: {release_revision: desc}, where: {release_name: {_eq: "${name}"}, release_namespace: {_eq: "${namespace}"}, release_revision: {_neq: ${currentRevision}}, release_cluster: {_eq: "${cluster}"}}) {
      release_revision
      app_version
      chart
      description
      status
      updated
    }
  }
`;

export const History = ({  release }) =>  {
    const { loading, error, data } = useQuery(GET_HISTORY(release.name, release.namespace, release.cluster, release.revision));
    const classes = useStyles();

    if (loading) return 'Loading...';
    if (error) return `Error! ${error.message}`;

    return (
        <Paper className={classes.root}>
        <Table className={classes.table} aria-label="simple table">
            <TableHead>
            <TableRow>
                <TableCell>Revision</TableCell>
                <TableCell align="right">Updated</TableCell>
                <TableCell align="right">Status</TableCell>
                <TableCell align="right">Chart</TableCell>
                <TableCell align="right">App Version</TableCell>
                <TableCell align="right">Description</TableCell>
                <TableCell align="right"></TableCell>
            </TableRow>
            </TableHead>
            <TableBody>
            {data.comodor_app_history.map(r => (
                <TableRow key={r.release_revision}>
                <TableCell component="th" scope="row">
                    {r.release_revision}
                </TableCell>
                <TableCell align="right">{r.updated}</TableCell>
                <TableCell align="right">{r.status}</TableCell>
                <TableCell align="right">{r.chart}</TableCell>
                <TableCell align="right">{r.app_version}</TableCell>
                <TableCell align="right">{r.description}</TableCell>
                <TableCell align="right"><Button variant="outlined">Diff</Button></TableCell>
                </TableRow>
            ))}
            </TableBody>
        </Table>
        </Paper>
    );
}