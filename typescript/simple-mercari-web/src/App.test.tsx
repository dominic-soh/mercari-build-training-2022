import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders Simple Mercari', () => {
  // Arrange and act
  render(<App />);
  const appTitle = screen.getByText(/Simple Mercari/i);
  
  // Assert
  expect(appTitle).toBeInTheDocument();
});

test('renders form component', async () => {
  // Arrange and act
  render(<App />);
  const appFormNameField = await screen.findByTestId(/name/i);
  const appFormCategoryField = await screen.findByTestId(/category/i);
  const appFormImageField = await screen.findByTestId(/image/i);
  const appFormSubmitButton = await screen.findByTestId(/image/i);
  
  // Assert
  expect(appFormNameField).toBeTruthy()
  expect(appFormCategoryField).toBeTruthy()
  expect(appFormImageField).toBeTruthy()
  expect(appFormSubmitButton).toBeTruthy()
});