import {backEndJson} from '../types/backEndJson'
import {Card} from 'react-bootstrap';

export const ArtCard = ({artWork}:{artWork: backEndJson}) => {
  return (
  <>
    <Card style={{width: '18rem'}} className='my-2'>
      {artWork&& <Card.Img variant='top' className='rounded h-200' src={artWork["image-url"]}/>} 
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