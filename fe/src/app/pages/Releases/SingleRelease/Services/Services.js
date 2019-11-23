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

export const Services = ({  release }) =>  {
    const { loading, error, data } = useQuery(GET_SERVICES(release.name, release.namespace, release.revision));
    const classes = useStyles();

    if (loading) return 'Loading...';
    if (error) return `Error! ${error.message}`;

    return (
        <Paper className={classes.root}>
        <Table className={classes.table} aria-label="simple table">
            <TableHead>
            <TableRow>
                <TableCell>Name</TableCell>
                <TableCell align="right">Type</TableCell>
                <TableCell align="right">Cluster IP</TableCell>
                <TableCell align="right">External IP</TableCell>
                <TableCell align="right">Port(s)</TableCell>
                <TableCell align="right">Created At</TableCell>
            </TableRow>
            </TableHead>
            <TableBody>
            {data.comodor_services.map(s => (
                <TableRow key={s.name}>
                <TableCell component="th" scope="row">
                    {s.name}
                </TableCell>
                <TableCell align="right">{s.type}</TableCell>
                <TableCell align="right">{s.cluster_ip}</TableCell>
                <TableCell align="right">{s.external_ip}</TableCell>
                <TableCell align="right">{s.ports}</TableCell>
                <TableCell align="right">{s.created_at}</TableCell>
                </TableRow>
            ))}
            </TableBody>
        </Table>
        </Paper>
    );
}