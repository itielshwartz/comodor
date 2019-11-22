import React from 'react';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

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
      <div style={{marginLeft: 700}} >
        <div>Releases</div>
        <list>
          {data.comodor_releases.map(r => 
            (<li>{`${r.cluster} - ${r.name}`}</li>)
          )}
        </list>
      </div>
  )
}