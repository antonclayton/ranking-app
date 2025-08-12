import React from 'react';
import { type Place as PlaceType } from '../../../../data/places';
import { Card, CardContent, Typography } from '@mui/material';

interface PlaceCardProps {
  place: PlaceType;
}

const PlaceCard = ({ place }: PlaceCardProps) => {
  return (
    <Card sx={{ height: '100%' }}>
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {place.name}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {place.tags.join(', ')}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default PlaceCard;