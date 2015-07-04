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
        public Vector2 Position { get; set; }

        public Vector2 Size { get; set; }
        

        public String Name { get; private set; }

        public Rectangle Bounds {
            get
            {
                return new Rectangle((int)this.Position.X, (int)this.Position.Y, (int)this.Size.X, (int)this.Size.Y);
            }
        }

        protected Texture2D Texture { get; private set; }
          
        public Sprite(MyGame g, String assetName, Rectangle size )
            : base(g)
        {
            this.Name = assetName;
            this.Position = size.Location.ToVector2();
            this.Size = new Vector2(size.Width, size.Height);
        }

        public void SetPosition(Point pos)
        {
            this.Position = pos.ToVector2();
            //this.SetPosition(pos.X, pos.Y);
        }

        public void SetPosition(float x, float y)
        {
            this.Position = new Vector2(x, y);
            //this.Bounds = new Rectangle( x, y, this.Bounds.Width, this.Bounds.Height);
        }

        protected override void LoadContent()
        {
            this.Texture = this.Game.Content.Load<Texture2D>(this.Name);
        }

        public override void Draw(GameTime gameTime)
        {
            MyGame.SpriteBatch.Draw(this.Texture, destinationRectangle: MyGame.Camera.TranslateToPixelsRect(this.Bounds));
        }
    }
}
