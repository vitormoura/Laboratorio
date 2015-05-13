using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    abstract class Sprite : DrawableGameComponent
    {
        public String Name { get; private set; }

        public Rectangle Bounds { get; private set; }

        protected Texture2D Texture { get; private set; }
          
        public Sprite(MyGame g, String assetName, Rectangle size )
            : base(g)
        {
            this.Name = assetName;
            this.Bounds = size;
        }

        public void SetPosition(Point pos)
        {
            this.Bounds = new Rectangle(pos.X, pos.Y, this.Bounds.Width, this.Bounds.Height);
        }

        protected override void LoadContent()
        {
            this.Texture = this.Game.Content.Load<Texture2D>(this.Name);
        }

        public override void Draw(GameTime gameTime)
        {
            SpriteBatch sb = this.Game.Services.GetService<SpriteBatch>();
            
            sb.Draw(this.Texture, destinationRectangle: this.Bounds);
        }
    }
}
