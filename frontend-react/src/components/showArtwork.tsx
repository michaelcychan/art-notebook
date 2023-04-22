import {Button, Card} from 'react-bootstrap';

import http from  '../http-common';
import { useEffect, useState } from 'react';

import {backEndJson} from '../types/backEndJson'

import {ArtCard} from './artCard'

export const ShowArtwork = () => {

  const emptyData:backEndJson = {
    "title": "Loading...",
    "image-url": "/sorry.gif",
    "short-description": "Loading...",
    "artist-title": [],
    "museum": "",
    "work-start": 0,
    "work-end": 0,
    "message": "error"
  }

  const [artWork, setArtWork] = useState<backEndJson>(emptyData)

  const fetchArtworkFromBackend = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-example-painting-Chicago")
      setArtWork(data)
    } catch {
      setArtWork(emptyData)
    }
  }

  useEffect(() => {
    fetchArtworkFromBackend()
  },[])

  return (
  <>
    <ArtCard artWork={artWork}/>
    <Button className='btn btn-info my-2' onClick={()=>fetchArtworkFromBackend()}>Reload</Button>
  </>
  )
}