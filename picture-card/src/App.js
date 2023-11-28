import { useState, useEffect } from 'react';
import Row from 'react-bootstrap/Row';

import PictureGrid from './component/PictureGrid';

import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

function App() {
  const [pictures, setPictures] = useState([]);
  useEffect(() => {
    PICTURE_DATA.map((picture) => {
      picture.image_tags = TAG_DATA.filter((tag) =>
        picture.tags ? picture.tags.includes(tag.id) : false
      );
    });
    setPictures(PICTURE_DATA);
  });
  return (
    <div className="App">
      <PictureGrid pictures={pictures} />
    </div>
  );
}

export default App;

const PICTURE_DATA = [
  {
    id: 1,
    tags: [1, 2, 3],
    name: 'Saint Mopierre',
    body: 'Saint Mopierre is a large city, known for being the birthplace of a music genre.',
    img_url: 'https://picsum.photos/id/11/1000',
  },
  {
    id: 2,
    tags: [1],
    name: 'Eulake',
    body: 'Eulake is a small town situated besides a lake. It is known for its mining heritage.',
    img_url: 'https://picsum.photos/id/11/1000',
  },
  {
    id: 3,
    tags: [2, 4],
    name: 'Prince Loeilles',
    body: 'Prince Loeilles is a large town, known for the battle of Prince Loeilles.',
    img_url: 'https://picsum.photos/id/11/1000',
  },
  {
    id: 4,
    tags: [2, 3, 4],
    name: 'North Warrines',
    body: 'North Warrines is a large town, known for being the birthplace of a music genre.',
    img_url: 'https://picsum.photos/id/11/1000',
  },
  {
    id: 5,
    items: [1, 2, 3],
    name: 'Sainttrois',
    body: 'Sainttrois is a large town named after Saint trois. It is known for the Sainttrois music festival.',
    img_url: 'https://picsum.photos/id/11/1000',
  },
  {
    id: 6,
    tags: [4],
    name: 'Grandenellakes',
    body: 'Grandenellakes is a large town situated besides a lake. It is known for its elaborate legends.',
    img_url: 'https://picsum.photos/id/200/800',
  },
];

const TAG_DATA = [
  {
    id: 1,
    name: 'Brinebeast',
    type: 'Earth',
  },
  {
    id: 2,
    name: 'Goolu',
    type: 'Air',
  },
  {
    id: 3,
    name: 'Macaronifeet',
    type: 'Fire',
  },
  {
    id: 4,
    name: 'Wispclaw',
    type: 'Water',
  },
];
