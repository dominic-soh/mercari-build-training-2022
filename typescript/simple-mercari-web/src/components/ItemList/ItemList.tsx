import React, { useEffect, useState } from 'react';

interface Item {
  Name: string;
  Category: string;
  Image_filename: string;
};

const server = process.env.API_URL || 'http://127.0.0.1:9000';
// const placeholderImage = process.env.PUBLIC_URL + '/logo192.png';

interface Prop {
  reload?: boolean;
  onLoadCompleted?: () => void;
}

export const ItemList: React.FC<Prop> = (props) => {
  const { reload = true, onLoadCompleted } = props;
  const [items, setItems] = useState<Item[]>([])
  const fetchItems = () => {
    fetch(server.concat('/items'),
      {
        method: 'GET',
        mode: 'cors',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
      })
      .then(response => response.json())
      .then(data => {
        console.log('GET success:', data);
        setItems(data.items);
        onLoadCompleted && onLoadCompleted();
      })
      .catch(error => {
        console.error('GET error:', error)
      })
  }

  useEffect(() => {
    if (reload) {
      fetchItems()
    }
  }, [reload]);
  

  return (
    <div className='ItemGrid'>
      {items.map((item, index) => {
        return (
          <div key={index} className='ItemList' data-testid={index}>
            {/* TODO: Task 1: Replace the placeholder image with the item image */}
            <img src={server + "/image/" + item.Image_filename} alt="dog"/>
            <p>
              <span>Name: {item.Name}</span>
              <br />
              <span>Category: {item.Category}</span>
            </p>
          </div>
        )
      })}
    </div>
  )
};
