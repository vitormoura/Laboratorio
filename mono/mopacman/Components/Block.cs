using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Block : Sprite
    {
        public MazeSection Section { get; private set; }
                
        public Block(MyGame g, MazeSection section)
            : base(g, "section.png", new Rectangle(0, 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH))
        {
            this.Section = section;
        }

        public override void Draw(GameTime gameTime)
        {
            Rectangle rectangleToDraw = new Rectangle(Constants.DEFAULT_BLOCK_WIDTH * (this.Section.HasCookie ? 0 : 1), 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH);

            MyGame.SpriteBatch.Draw(this.Texture, destinationRectangle: MyGame.Camera.TranslateToPixelsRect(this.Bounds), sourceRectangle: rectangleToDraw);
        }
    }
}
