using Microsoft.Xna.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Animations
{
    class BlinkAnimation<T>
        where T : DrawableGameComponent
    {
        public T Item { get; private set; }

        public BlinkAnimation( T item )
        {
            this.elapsed = 0.2f;
            this.Item = item;
        }

        public void Update(GameTime gameTime)
        {
            if (this.elapsed <= 0.0f)
            {
                this.elapsed = 0.2f;
                this.Item.Visible = !this.Item.Visible;
            }
            else
                this.elapsed -= (float)gameTime.ElapsedGameTime.TotalSeconds;
        }

        private float elapsed;
    }
}
