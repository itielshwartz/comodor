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



const GET_DEPLOYMENTS = (name, namespace, revision) => gql`
  {
    comodor_deployments(where: {unique_release_name: {_eq: "${name}"}, unique_release_namespace: {_eq: "${namespace}"}, unique_release_revision: {_eq: ${revision}}}) {
      available
      created_at
      name
      ready
      total
    }
  }
`;

export const Deployments = ({  release }) =>  {
    const { loading, error, data } = useQuery(GET_DEPLOYMENTS(release.name, release.namespace, release.revision));
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
                <TableCell align="right">Available</TableCell>
                <TableCell align="right">Created At</TableCell>
            </TableRow>
            </TableHead>
            <TableBody>
            {data.comodor_deployments.map(d => (
                <TableRow key={d.name}>
                <TableCell component="th" scope="row">
                    {d.name}
                </TableCell>
                <TableCell align="right">{`${d.ready}/${d.total}`}</TableCell>
                <TableCell align="right">{d.available}</TableCell>
                <TableCell align="right">{d.created_at}</TableCell>
                </TableRow>
            ))}
            </TableBody>
        </Table>
        </Paper>
    );
}