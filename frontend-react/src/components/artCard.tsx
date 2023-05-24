import {backEndJson} from '../types/backEndJson'
import {Card} from 'react-bootstrap';
import {ShowTagSpan} from './cardTag';

export const ArtCard = ({artWork}:{artWork: backEndJson}) => {

  const showDesc = (text:string) => {
    if (text.length > 50) {
      return text.slice(0, 50) + "...";
    } else if (text.length > 0) {
      return text
    } else {
      return "[No description to display]"
    }
  }

  return (
  <>
    <Card style={{width: '18rem'}} className='my-2 mx-3 rounded-3 artcard'>
      <Card.Header className='h3'>{artWork.museum.length > 0 ? artWork.museum : "Museum Name"}</Card.Header>
      {artWork&& <Card.Img variant='top' className='rounded h-200 mx-1 my-1' src={artWork["image-url"]}/>} 
      <Card.Body>
        <Card.Title>{artWork ? artWork.title : "This is the Title"} </Card.Title>
        <Card.Text>
          <div className='description-container'>
            {showDesc(artWork['short-description']) }
          </div>
        </Card.Text>
      </Card.Body>
      <Card.Footer>
        {(artWork.Tags && artWork.Tags.length > 0) &&
        <>
          <div>Tags:</div>
          <ShowTagSpan tags={artWork.Tags}/>
        </>
        }
        {(artWork.note && artWork.note.length > 0) &&
          <>
            <div>Note:</div>
            {artWork.note}
          </>
        }
      </Card.Footer>
    </Card>
  </>
  )
}