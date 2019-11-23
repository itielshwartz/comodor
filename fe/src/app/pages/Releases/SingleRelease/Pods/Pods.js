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


const useStyles = makeStyles({
    root: {
      width: '100%',
      overflowX: 'auto',
    },
    table: {
      minWidth: 650,
    },
  });



const GET_PODS = (name, namespace, revision) => gql`
  {
    comodor_pods(where: {unique_release_name: {_eq: "${name}"}, unique_release_namespace: {_eq: "${namespace}"}, unique_release_revision: {_eq: ${revision}}}) {
      name
      ready
      total
      restarts
      status
      created_at
    }
  }
`;

export const Pods = ({  release }) =>  {
    const { loading, error, data } = useQuery(GET_PODS(release.name, release.namespace, release.revision));
    const classes = useStyles();

    if (loading) return 'Loading...';
    if (error) return `Error! ${error.message}`;

    return (
        <Paper className={classes.root}>
        <Table className={classes.table} aria-label="simple table">
            <TableHead>
            <TableRow>
                <TableCell>Name</TableCell>
                <TableCell align="right">Ready</TableCell>
                <TableCell align="right">Status</TableCell>
                <TableCell align="right">Restarts</TableCell>
                <TableCell align="right">Created At</TableCell>
            </TableRow>
            </TableHead>
            <TableBody>
            {data.comodor_pods.map(p => (
                <TableRow key={p.name}>
                <TableCell component="th" scope="row">
                    {p.name}
                </TableCell>
                <TableCell align="right">{`${p.ready}/${p.total}`}</TableCell>
                <TableCell align="right">{p.status}</TableCell>
                <TableCell align="right">{p.restarts}</TableCell>
                <TableCell align="right">{p.created_at}</TableCell>
                </TableRow>
            ))}
            </TableBody>
        </Table>
        </Paper>
    );
}