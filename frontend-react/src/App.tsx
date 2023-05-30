import React from 'react';
import './App.css';
import { ShowArtwork } from './components/showArtwork';

import 'bootstrap/dist/css/bootstrap.min.css';
import { ShowSavedWork } from './components/showSaved';

function App() {
  return (
    <div className="App">
      <div className="header">
        <div id="title">
          <h1 id='title' className='display-1'>
            Art Notebook
          </h1>
        </div>
      </div>
        <ShowArtwork />

      <hr />
      <ShowSavedWork />
    </div>
  );
}

export default App;
