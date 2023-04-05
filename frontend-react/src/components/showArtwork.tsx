import {Card} from 'react-bootstrap';

import http from  '../http-common';
import { useEffect, useState } from 'react';

type backEndJson = {
  "title": string,
  "short-description": string,
  "artist-title": string[],
  "image-url": string,
  "museum": string,
  "work-start":number,
  "work-end":number,
  "message": string,
}

export const ShowArtwork = () => {

  const emptyData:backEndJson = {
    "title": "Sorry! Cannot load!",
    "image-url": "/sorry.gif",
    "short-description": "Cannot load!",
    "artist-title": [],
    "museum": "",
    "work-start": 0,
    "work-end": 0,
    "message": "error"
  }

  const [artWork, setArtWork] = useState<backEndJson>()

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
    <Card style={{width: '18rem'}}>
      {artWork&& <Card.Img variant='top' src={artWork["image-url"]}/>} 
      <Card.Body>
        <Card.Title>{artWork ? artWork.title : "This is the Title"} </Card.Title>
        <Card.Text>{artWork ? artWork['short-description'] : "This is some card text"}</Card.Text>
      </Card.Body>
      <Card.Footer>
        {artWork ? artWork.museum : "failed to load"}
      </Card.Footer>

    </Card>
  </>
  )
}