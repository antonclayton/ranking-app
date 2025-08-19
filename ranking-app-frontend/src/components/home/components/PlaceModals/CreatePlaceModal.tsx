import React, { useState } from 'react';
import { createPlace } from '../../../../utils/api';
import { type Place } from '../../../../types/placeTypes';
import { Dialog, DialogTitle, DialogContent, DialogActions, Button, TextField, Stack } from '@mui/material';

interface CreatePlaceModalProps {
  open: boolean;
  onClose: () => void;
  onCreated: (place: Place) => void;
}

const CreatePlaceModal: React.FC<CreatePlaceModalProps> = ({ open, onClose, onCreated }) => {
  const [name, setName] = useState('');
  const [tagsInput, setTagsInput] = useState('');
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const reset = () => {
    setName('');
    setTagsInput('');
    setError(null);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setSubmitting(true);
    setError(null);
    try {
      const tags = tagsInput
        .split(',')
        .map((t) => t.trim())
        .filter(Boolean);
      const created = await createPlace({ name: name.trim(), tags });
      onCreated(created);
      reset();
      onClose();
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Failed to create place';
      setError(message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="sm">
      <DialogTitle>Create Place</DialogTitle>
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
              placeholder="pizza, italian, casual"
            />
            {error && <div style={{ color: 'var(--error, #f44336)' }}>{error}</div>}
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose} disabled={submitting}>Cancel</Button>
          <Button type="submit" variant="contained" disabled={submitting || name.trim().length === 0}>
            {submitting ? 'Creating...' : 'Create'}
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
};

export default CreatePlaceModal;