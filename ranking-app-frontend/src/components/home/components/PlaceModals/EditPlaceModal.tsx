import React, { useState, useEffect } from 'react';
import { Dialog, DialogTitle, DialogContent, DialogActions, Button, TextField, Stack, Typography } from '@mui/material';
import { type Place } from '../../../../types/placeTypes';
import { updatePlace } from '../../../../utils/api';

interface EditPlaceModalProps {
  open: boolean;
  onClose: () => void;
  place: Place; // initial values
  onSaved: (place: Place) => void; // callback with updated place
}

const EditPlaceModal: React.FC<EditPlaceModalProps> = ({ open, onClose, place, onSaved }) => {
  const [name, setName] = useState(place.name);
  const [tagsInput, setTagsInput] = useState(place.tags.join(', '));
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Reset form when the dialog opens or the place changes
  useEffect(() => {
    if (open) {
      setName(place.name);
      setTagsInput(place.tags.join(', '));
      setError(null);
    }
  }, [open, place]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setSaving(true);
    setError(null);
    try {
      const tags = tagsInput.split(',').map((t) => t.trim()).filter(Boolean);
      const updated = await updatePlace(place.id, { name: name.trim(), tags });
      onSaved(updated);
      onClose();
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update place');
    } finally {
      setSaving(false);
    }
  };

  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="sm">
      <DialogTitle>Edit Place</DialogTitle>
      <form onSubmit={handleSubmit}>
        <DialogContent>
          <Stack spacing={2}>
            <TextField
              label="Name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
              autoFocus
            />
            <TextField
              label="Tags (comma-separated)"
              value={tagsInput}
              onChange={(e) => setTagsInput(e.target.value)}
            />
            {error && (
              <Typography color="error" variant="body2">
                {error}
              </Typography>
            )}
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose} disabled={saving}>
            Cancel
          </Button>
          <Button type="submit" variant="contained" disabled={saving || name.trim().length === 0}>
            {saving ? 'Saving...' : 'Save'}
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
};

export default EditPlaceModal;