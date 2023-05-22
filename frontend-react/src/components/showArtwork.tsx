import {Button} from 'react-bootstrap';

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
    "source-id": "",
    "message": "error",
    "Tags": [],
    "note": ""
  }

  const [artWork, setArtWork] = useState<backEndJson>(emptyData)

  const fetchArtworkFromArtChicago = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-example-painting-Chicago")
      setArtWork(data)
    } catch {
      setArtWork(emptyData)
    }
  }

  const fetchArtworkFromMetroMArt = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-example-painting-Metro")
      setArtWork(data)
    } catch {
      setArtWork(emptyData)
    }
  }

  const fetchNpmArtwork = async () => {
    try {
      const {data} = await http.get<backEndJson>("/get-painting-npm")
      setArtWork(data)
    } catch {
      setArtWork(emptyData)
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