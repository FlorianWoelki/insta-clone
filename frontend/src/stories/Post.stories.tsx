import React from 'react';
import Post from '../components/Post';

export default {
  title: 'Components/Post',
  component: Post,
};

export const Default = () => (
  <div className="max-w-md">
    <Post></Post>
  </div>
);
