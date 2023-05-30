import {Button} from 'react-bootstrap';

import http from  '../http-common';
import { useEffect, useRef, useState } from 'react';

import {backEndJson} from '../types/backEndJson'

import {ArtCard} from './artCard'
import { AxiosError } from 'axios';

export const ShowArtwork = () => {

  const errorGettingData:backEndJson = {
    "title": "Loading...",
    "image-url": "/sorry.gif",
    "short-description": "",
    "artist-title": [],
    "museum": "",
    "work-start": 0,
    "work-end": 0,
    "source-id": "",
    "tags": [],
    "note": ""
  }

  const [artWork, setArtWork] = useState<backEndJson>({
    "title": "Loading...",
    "image-url": "/sorry.gif",
    "short-description": "Loading...",
    "artist-title": [],
    "museum": "Loading...",
    "work-start": 0,
    "work-end": 0,
    "source-id": "",
    "tags": [],
    "note": ""
  })

  const fetchIsCancelled = useRef(false);

  const fetchArtwork = async (museum:String) => {
    const museumMap: {[id:string]:string} = {
      "chicago": "get-example-painting-Chicago",
      "npm": "get-painting-npm",
      "metro": "get-example-painting-Metro"
    }
    try {
      const endpoint = "/" + museumMap[`${museum}`];
      const {data} = await http.get<backEndJson>(endpoint);
      
      if (!fetchIsCancelled.current) {
        setArtWork(data)
      }
    } catch (err) {
      if (err instanceof AxiosError) {
        console.error("cannot connect to server");
      }
      setArtWork(errorGettingData);
    }
  }

  useEffect(() => {
    fetchArtwork("chicago");
    return () => {
      fetchIsCancelled.current = true;
    }
  },[])

  return (
  <>
    <div className="random-container">
      <ArtCard artWork={artWork}/>
    </div>
    <Button className='btn btn-info my-2 mx-2' onClick={()=>fetchArtwork("chicago")}>Art Institute of Chicago</Button>
    <Button className='btn btn-info my-2 mx-2' onClick={()=>fetchArtwork("npm")}>National Palace Museum 台灣故宮</Button>
    <Button className='btn btn-info my-2 mx-2' onClick={()=>fetchArtwork("metro")}>Metropolitan Museum of Art</Button>
    
  </>
  )
}