import React from 'react';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const PictureGrid = ({pictures}) => {
  return (
    <>
      <Row className="mt-5">
        <Col className="col-lg-6 offset-lg-3 col-md-8 offset-md-2 col-12">
          <h3>Picture Grid Display</h3>
        </Col>
      </Row>
    </>
  );
};

export default PictureGrid;
