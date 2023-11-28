import { useState, useEffect } from 'react';
import { getImages, getImageTags } from './services/imageService';

import PictureGrid from './components/PictureGrid';

import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

function App() {
  const [pictures, setPictures] = useState([]);

  useEffect(() => {
    let isMounted = true;

    (async () => {
      try {
        const pictures = await getImages();
        const tags = await getImageTags();
        pictures.map((picture) => {
          picture.image_tags = tags.filter((tag) =>
            picture.tags ? picture.tags.includes(tag.id) : false
          );
        });
        setPictures(pictures);
      } catch (err) {
        alert(`failed to load images: ${err}`);
      }
    })();

    // Cleanup callback as the component unmounts.
    return () => {
      isMounted = false;
    };
  }, []);
  return (
    <div className="App">
      <PictureGrid pictures={pictures} />
    </div>
  );
}

export default App;
