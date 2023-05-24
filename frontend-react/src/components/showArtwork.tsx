import {Button} from 'react-bootstrap';

import http from  '../http-common';
import { useEffect, useState } from 'react';

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
    "message": "error",
    "Tags": [],
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
    "message": "loading",
    "Tags": [],
    "note": ""
  })

  const fetchArtworkFromArtChicago = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-example-painting-Chicago")
      setArtWork(data)
    } catch (err:unknown) {
      if (err instanceof AxiosError) {
        console.error("cannot connect to server");
      }
      setArtWork(errorGettingData);
    }
  }

  const fetchArtworkFromMetroMArt = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-example-painting-Metro")
      setArtWork(data)
    } catch {
      setArtWork(errorGettingData)
    }
  }

  const fetchNpmArtwork = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-painting-npm")
      setArtWork(data)
    } catch {
      setArtWork(errorGettingData)
    }
  }

  useEffect(() => {
    fetchArtworkFromArtChicago()
  },[])

  return (
  <>
    <ArtCard artWork={artWork}/>
    <Button className='btn btn-info my-2 mx-2' onClick={()=>fetchArtworkFromArtChicago()}>Art Institute of Chicago</Button>
    <Button className='btn btn-info my-2 mx-2' onClick={()=>fetchNpmArtwork()}>National Palace Museum 台灣故宮</Button>
    <Button className='btn btn-info my-2 mx-2' onClick={()=>fetchArtworkFromMetroMArt()}>Metropolitan Museum of Art</Button>
    
  </>
  )
}