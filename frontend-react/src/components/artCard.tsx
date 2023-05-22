import {backEndJson} from '../types/backEndJson'
import {Card} from 'react-bootstrap';
import {ShowTagSpan} from './cardTag';

export const ArtCard = ({artWork}:{artWork: backEndJson}) => {

  const shortenText = (text:string) => {
    if (text.length <= 50) {
      return text
    } else {
      return text.slice(0, 50) + "...";
    }
  }

  return (
  <>
    <Card style={{width: '18rem'}} className='my-2'>
      {artWork&& <Card.Img variant='top' className='rounded h-200' src={artWork["image-url"]}/>} 
      <Card.Header>{artWork ? artWork.museum : "failed to load"}</Card.Header>
      <Card.Body>
        <Card.Title>{artWork ? artWork.title : "This is the Title"} </Card.Title>
        <Card.Text>{artWork ? shortenText(artWork['short-description']) : "This is some card text"}</Card.Text>
      </Card.Body>
      <Card.Footer>
        {artWork.Tags &&
        <>
          <div>Tags:</div>
          <ShowTagSpan tags={artWork.Tags}/>
        </>
        }
      </Card.Footer>
    </Card>
  </>
  )
}