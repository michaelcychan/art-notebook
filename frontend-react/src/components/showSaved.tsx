import { useState } from 'react';
import http from  '../http-common';

import {backEndJson} from '../types/backEndJson';
import {ArtCard} from './artCard';

type savedData = {
  data: backEndJson[]
}

export const ShowSavedWork = () => {

  const [savedData, setSavedData] = useState<backEndJson[]>([]);

  const username = "michael";

  const fetchSavedWork = async (username: String) => {
    try {
      const endpointUri = "/get-saved-data" + "?username=" + username
      const {data} = await http.get<savedData>(endpointUri)
      if (data.data.length > 0) {
        setSavedData(data.data);
      }
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <>
      <button onClick={() => fetchSavedWork(username)}>click to retrieve saved data</button>

      <div className="card-container">

        {savedData.length > 0 &&
          savedData.map(artData => <ArtCard artWork={artData} key={artData.museum + artData['source-id']}/>
          )
        }
      </div>
    </>
  )
}