import http from  '../http-common';

import {backEndJson} from '../types/backEndJson'

type savedData = {
  data: backEndJson[]
}

export const ShowSavedWork = () => {

  const username = "michael";

  const fetchSavedWork = async (username: String) => {
    try {
      const endpointUri = "/get-saved-data" + "?username=" + username
      console.log(endpointUri)
      const {data} = await http.get<savedData>(endpointUri)
      console.table(data);
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <button onClick={() => fetchSavedWork(username)}>click to retrieve saved data</button>
  )

}