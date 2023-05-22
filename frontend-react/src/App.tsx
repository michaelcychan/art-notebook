import React from 'react';
import './App.css';
import { ShowArtwork } from './components/showArtwork';

import 'bootstrap/dist/css/bootstrap.min.css';
import { ShowSavedWork } from './components/showSaved';

function App() {
  return (
    <div className="App">
      <ShowArtwork />
      <hr />
      <ShowSavedWork />
    </div>
  );
}

export default App;
