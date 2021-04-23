import React from 'react';
import Button from '../components/Button';

export default {
  title: 'UI/Button',
  component: Button,
};

interface PlaygroundArgs {
  content: string;
}

export const Default = () => <Button>Default Button</Button>;
export const Playground = ({ content }: PlaygroundArgs) => (
  <Button>{content}</Button>
);
Playground.args = { content: 'Playground Button' };
