import React from 'react';
import { places, type Place as PlaceType } from '../../../../data/places';
import PlaceCard from './PlaceCard';
import styles from './PlaceList.module.css';

interface PlaceListProps {
  searchQuery: string;
}

const PlaceList: React.FC<PlaceListProps> = ({ searchQuery }) => {
  const filteredPlaces = places.filter(
    (place) =>
      place.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      place.types.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className={styles.placeListContainer}>
      <div className={styles.placeList}>
        {filteredPlaces.length > 0 ? (
          filteredPlaces.map((place: PlaceType) => (
            <PlaceCard key={place.id} place={place} />
          ))
        ) : (
          <p>No restaurants found.</p>
        )}
      </div>
    </div>
  );
};

export default PlaceList;