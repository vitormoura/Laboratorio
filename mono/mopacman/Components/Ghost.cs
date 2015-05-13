using Microsoft.Xna.Framework;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Ghost : Player
    {
        public Ghost(MyGame g)
            : base(g, "ghost.png", new Rectangle(0, 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH))
        {
        }
    }
}
