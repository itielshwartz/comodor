import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import StatusIndicator from '../../components/StatusIndicator';
import FilterBar from '../../components/FilterBar';
import '../../assets/css/releases.css';

const GET_RELEASES = gql`
  {
    comodor_releases {
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

export const Releases = () =>  {

  const { loading, error, data } = useQuery(GET_RELEASES);

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  return (
      <div style={{padding: 24}} >
        <h1>Releases</h1>
        <FilterBar />
        <div className="card-list box-h">
          {data.comodor_releases.map(r => 
            (<Card key={r.row_id}>
              <CardContent>
                <div className="release-cluster">{r.cluster}</div>
                <div className="release-name">{r.name}</div>
                <StatusIndicator status={r.status}/>
                <div className="release-namespace">{r.namespace}</div>
              </CardContent>
            </Card>)
          )}
        </div>
      </div>
  )
}