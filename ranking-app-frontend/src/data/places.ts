export interface Place {
  id: number;
  name: string;
  tags: string[];
}

export const places: Place[] = [
  { id: 1, name: 'The Golden Spoon', tags: ['Italian', 'Japanese'] },
  { id: 2, name: 'Sushi Palace', tags: ['Japanese'] },
  { id: 3, name: 'Taco Fiesta', tags: ['Mexican'] },
  { id: 4, name: 'Burger Barn', tags: ['American'] },
  { id: 5, name: 'Curry House', tags: ['Indian'] },
  { id: 6, name: 'Pizza Planet', tags: ['Italian'] },
  { id: 7, name: 'Noodle Nirvana', tags: ['Thai'] },
  { id: 8, name: 'Steakhouse Supreme', tags: ['American'] },
];
