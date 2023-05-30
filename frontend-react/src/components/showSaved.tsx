import { FormEvent, useRef, useState } from 'react';
import http from  '../http-common';

import {backEndJson} from '../types/backEndJson';
import {ArtCard} from './artCard';

type savedData = {
  data: backEndJson[]
}

export const ShowSavedWork = () => {

  const [savedData, setSavedData] = useState<backEndJson[]>([]);

  const usernameRef = useRef<HTMLInputElement | null>(null);

  const handleSubmit = async (e:FormEvent) => {
    e.preventDefault();
    if (!usernameRef.current) {
      return
    }
    try {
      const endpointUri = "/get-saved-data?username=" + usernameRef.current.value;
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
      <form onSubmit={handleSubmit}>
        <label htmlFor="username-notebook">Username:</label>
        <input id='username-notebook' name='username-notebook' ref={usernameRef} />
        <button type='submit'>Check notebook</button>
      </form>
      <div className="card-container container-fluid">
        <div className="row">
          {savedData.length > 0 &&
            savedData.map(artData =>{ 
              return (
                <ArtCard artWork={artData} key={artData.museum + artData['source-id']}/>
              ) 
            }
            )
          }
        </div>
      </div>
    </>
  )
}