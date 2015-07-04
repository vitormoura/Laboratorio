using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Text : DrawableGameComponent
    {
        public String Message { get; set; }

        public Vector2 Position { get; set; }

        public Color Color { get; set; }
        
        public Text(MyGame g, SpriteFont font, String msg)
            : base(g)
        {
            this.Message = msg;
            this.Color = Color.White;

            this.font = font;
        }

        public override void Update(GameTime gameTime)
        {
            base.Update(gameTime);
        }

        public override void Draw(GameTime gameTime)
        {
            MyGame.SpriteBatch.DrawString(this.font, this.Message, MyGame.Camera.TranslateToPixels(this.Position).ToVector2(), this.Color);
        }

        private SpriteFont font;
    }
}
