export interface Place {
  id: number;
  name: string;
  types: string;
}

export const places: Place[] = [
  { id: 1, name: 'The Golden Spoon', types: 'Italian' },
  { id: 2, name: 'Sushi Palace', types: 'Japanese' },
  { id: 3, name: 'Taco Fiesta', types: 'Mexican' },
  { id: 4, name: 'Burger Barn', types: 'American' },
  { id: 5, name: 'Curry House', types: 'Indian' },
  { id: 6, name: 'Pizza Planet', types: 'Italian' },
  { id: 7, name: 'Noodle Nirvana', types: 'Thai' },
  { id: 8, name: 'Steakhouse Supreme', types: 'American' },
];
