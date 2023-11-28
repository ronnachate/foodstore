import React from 'react';
import Card from 'react-bootstrap/Card';
import Col from 'react-bootstrap/Col';
import Badge from 'react-bootstrap/Badge';

import './picture-card.css';

const PictureCard = ({ picture }) => {
  return (
    <Col className="col-xl-3 col-lg-6 col-md-6 col-12 pb-4">
      <Card className="picture-card shadow">
        <Card.Body>
          <Card.Title>
            <Card.Img variant="top" className="mt-3" src={picture.img_url} />
            <div className="picture-title">
              <span className="mt-3">{picture.name}</span>
            </div>
          </Card.Title>
          <Card.Text className="mt-4 pt-1">{picture.body}</Card.Text>
        </Card.Body>
        <div className="d-flex gap-2 footer">
          {picture.image_tags.map((tag) => (
            <Badge key={tag.id}>{tag.name}</Badge>
          ))}
        </div>
      </Card>
    </Col>
  );
};

export default PictureCard;
