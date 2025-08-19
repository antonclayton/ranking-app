import React, { useState } from 'react';
import { type Place as PlaceType } from '../../../../types/placeTypes';
import { Card, CardContent, Typography, CardActions, Button } from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';
import EditPlaceModal from '../../components/PlaceModals/EditPlaceModal';

interface PlaceCardProps {
  place: PlaceType;
  onUpdated?: (place: PlaceType) => void;
}

const PlaceCard = ({ place, onUpdated }: PlaceCardProps) => {
  const [open, setOpen] = useState(false);

  const handleOpen = () => {
    setOpen(true);
  };
  const handleClose = () => setOpen(false);

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
      <CardActions sx={{ justifyContent: 'flex-end' }}>
        <Button size="small" startIcon={<EditIcon />} onClick={handleOpen}>
          Edit
        </Button>
      </CardActions>

      <EditPlaceModal
        open={open}
        onClose={handleClose}
        place={place}
        onSaved={(updated) => {
          onUpdated?.(updated);
          setOpen(false);
        }}
      />
    </Card>
  );
};

export default PlaceCard;