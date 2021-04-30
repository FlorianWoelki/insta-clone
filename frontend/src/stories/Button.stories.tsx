import React from 'react';
import Button from '../components/Button';
import { ButtonType } from '../components/Button';

export default {
  title: 'UI/Button',
  component: Button,
};

interface PlaygroundArgs {
  content: string;
  type: ButtonType;
}

export const Default = () => <Button>Default Button</Button>;
export const PrimaryButton = () => (
  <Button type="primary">Primary Button</Button>
);
export const SecondaryButton = () => (
  <Button type="secondary">Secondary Button</Button>
);
export const Playground = ({ content, type }: PlaygroundArgs) => (
  <Button type={type}>{content}</Button>
);
Playground.args = { content: 'Playground Button', type: 'primary' };
