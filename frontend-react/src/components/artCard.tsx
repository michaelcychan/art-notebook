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
    <Card className='my-2 mx-3 rounded-3 artcard' style={{minHeight:"75px"}}>
      <Card.Header className='h3'>{artWork.museum.length > 0 ? artWork.museum : "Museum Name"}</Card.Header>
      <div className="row no-gutters" >
        <div className="col-md-4">

        {artWork&& <Card.Img className='rounded h-200 mx-1 my-1' src={artWork["image-url"]}/>} 
        </div>
        <div className="col-md-8">
          <Card.Body>
            <Card.Title>{ artWork.title} </Card.Title>
            <Card.Text>
                {showDesc(artWork['short-description']) }
            </Card.Text>
          </Card.Body>
          <Card.Footer>
            {(artWork.tags && artWork.tags.length > 0) &&
            <>
              <div>Tags:</div>
              <ShowTagSpan tags={artWork.tags}/>
            </>
            }
            {(artWork.note && artWork.note.length > 0) &&
              <>
                <div>Note:</div>
                {artWork.note}
              </>
            }
          </Card.Footer>
        </div>
      </div>
    </Card>
  </>
  )
}