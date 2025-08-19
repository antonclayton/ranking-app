import styles from './PlaceList.module.css';
import { type Place as PlaceType } from '../../../../types/placeTypes';
import { useState, useEffect } from 'react';
import PlaceCard from './PlaceCard';
import { getPlaces } from '../../../../utils/api';

interface PlaceListProps {
  searchQuery: string;
  refreshToken: number;
  onEdited: (place: PlaceType) => void;
}

const PlaceList: React.FC<PlaceListProps> = ({ searchQuery, refreshToken, onEdited }) => {
  const [places, setPlaces] = useState<PlaceType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        setError(null);
        const data = await getPlaces();
        setPlaces(data);
      } catch (err: unknown) {
        setError(err instanceof Error ? err.message : 'Unknown error occurred');
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [refreshToken]);

  const handleUpdated = (updated: PlaceType) => {
    // Update locally for immediate UI feedback
    setPlaces((prev) => prev.map((p) => (p.id === updated.id ? updated : p)));
    // Let parent (Home) know so it can optionally refetch via refreshToken
    onEdited(updated);
  };

  const filteredPlaces = places.filter(
    (place) =>
      place.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      place.tags.join(',').toLowerCase().includes(searchQuery.toLowerCase())
  );

  if (loading) {
    return (
      <div className={styles.placeListContainer}>
        <p className={styles.placeList}>Loading...</p>
      </div>
    );
  }
  if (error) {
    return (
      <div className={styles.placeListContainer}>
        <p className={styles.placeList}>Error: {error}</p>
      </div>
    );
  }

  return (
    <div className={styles.placeListContainer}>
      <div className={styles.placeList}>
        {filteredPlaces.length > 0 ? (
          filteredPlaces.map((place: PlaceType) => (
            <PlaceCard key={place.id} place={place} onUpdated={handleUpdated} />
          ))
        ) : (
          <p>No restaurants found.</p>
        )}
      </div>
    </div>
  );
};

export default PlaceList;