import React, { useState, useEffect } from 'react';
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
  const [ allReleases, setAllReleases ] = useState([])
  const [ visibleReleases, setVisibleReleases ] = useState([])

  function filterChanged(filters) {
    var filtered = allReleases;

    filtered = filtered.filter((item) => {
      var isValid = true;

      for (var i in filters) {
        let filter = filters[i].value.split("_"),
          field = filter[0],
          value = filter[1];
        
        isValid = isValid && item[field] === value
      }

      return isValid;
    });

    setVisibleReleases(filtered);
  }

  function cardClicked(id) {
    window.location.href = "releases/" + id;
  }

  useEffect(() => {
    if (data) {
      setVisibleReleases(data.comodor_releases)
      setAllReleases(data.comodor_releases)
    }
  }, [data])

  if (loading) return 'Loading...';
  if (error) return `Error! ${error.message}`;

  return (
      <div style={{padding: 24}} >
        <h1>Releases</h1>
        <FilterBar onChange={filterChanged} />
        <div className="card-list box-h">
          {visibleReleases.map(r => 
            (<Card key={r.row_id} onClick={cardClicked.bind(this, r.row_id)}>
              <CardContent>
                <div className="release-cluster">{r.cluster}</div>
                <div className="release-name">{r.name}</div>
                <StatusIndicator status={r.status}/>
                <div className="release-revision">#{r.revision}</div>
                <div className="release-namespace">{r.namespace}</div>
              </CardContent>
            </Card>)
          )}
        </div>
      </div>
  )
}