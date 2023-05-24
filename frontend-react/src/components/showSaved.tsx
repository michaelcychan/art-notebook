import { ChangeEvent, useState } from 'react';
import http from  '../http-common';

import {backEndJson} from '../types/backEndJson';
import {ArtCard} from './artCard';

type savedData = {
  data: backEndJson[]
}

export const ShowSavedWork = () => {

  const [savedData, setSavedData] = useState<backEndJson[]>([]);

  const [username, setUsername] = useState("");

  const handleUsernameChange = (username:string) => {
    setUsername(username);
  }

  const fetchSavedWork = async (username: String) => {
    try {
      const endpointUri = "/get-saved-data?username=" + username
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
      <label htmlFor="username-notebook">Username:</label>
      <input id='username-notebook' type="text" value={username} onChange={(event:ChangeEvent<HTMLInputElement>) => handleUsernameChange(event.target.value)}/>
      <button onClick={() => fetchSavedWork(username)}>Check notebook</button>
      <div className="card-container container-fluid">
        <div className="row">
          {savedData.length > 0 &&
            savedData.map(artData =>{ 
              return (
                <div className='col justify-content-center col-xl-3 col-lg-4 col-sm-12'  key={artData.museum + artData['source-id']}>
                <ArtCard artWork={artData}/>
                </div>
              ) 
            }
            )
          }
        </div>
      </div>
    </>
  )
}