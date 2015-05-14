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
        public Block(MyGame g)
            : base(g, "block.png", new Rectangle(0, 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH))
        {
        }
    }
}
