import React from 'react';
import Avatar from '../components/Avatar';
import Button from '../components/Button';
import Icon from '../components/Icon';
import Post from '../components/Post';
import Search from '../components/Search';

const Home = (): JSX.Element => {
  return (
    <React.Fragment>
      <div className="flex items-center justify-between">
        <Search />

        <div className="flex items-center space-x-6">
          <a href="#" className="relative text-gray-600">
            <Icon name="bell" className="w-5 h-5" />
            <div className="absolute top-0 right-0 w-1 h-1 -mt-1 -mr-1 bg-red-600 rounded-full" />
          </a>
          <a href="#" className="relative text-gray-600">
            <Icon name="mail" className="w-5 h-5" />
            <div className="absolute top-0 right-0 w-1 h-1 -mt-1 -mr-1 bg-red-600 rounded-full" />
          </a>
          <Button>
            <Icon name="plus-circle" className="w-6 h-6" />
            <span>Add photo</span>
          </Button>
        </div>
      </div>

      <div className="mt-10">
        <div className="flex items-center justify-between">
          <h1 className="text-3xl font-bold text-gray-800">Stories</h1>
          <div className="flex items-center space-x-3 cursor-pointer">
            <Icon name="play" className="w-6 h-6 text-gray-300" />
            <p className="text-gray-800">Watch all</p>
          </div>
        </div>
        <div className="flex items-center mt-6 space-x-4">
          <Avatar noActiveStories>
            <div className="absolute inset-0 flex items-center justify-center w-full h-full bg-blue-600 bg-opacity-50 rounded-full">
              <Icon name="plus" className="w-6 h-6 text-white" />
            </div>
          </Avatar>
          <Avatar />
          <Avatar />
          <Avatar noActiveStories />
        </div>
      </div>

      <div className="mt-10">
        <h1 className="text-3xl font-bold text-gray-800">Feed</h1>
        <div className="grid grid-cols-3 gap-6 mt-6">
          <Post />
          <Post />
          <Post />
          <Post />
          <Post />
        </div>
      </div>
    </React.Fragment>
  );
};

export default Home;
