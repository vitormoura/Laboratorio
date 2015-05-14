using Microsoft.Xna.Framework;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class MazeSection
    {
        public Vector2 ID { get; private set; }

        public Boolean Allowed { get; set; }

        public Boolean Checkpoint { get; set; }
                
        public MazeSection N { get; set; }

        public MazeSection S { get; set; }

        public MazeSection W { get; set; }

        public MazeSection E { get; set; }

        public MazeSection(Int32 x, Int32 y)
        {
            this.ID = new Vector2(x,y);
        }

        public MazeSection Get(EnumDirections d)
        {
            switch (d)
            {
                case EnumDirections.North:
                    return this.N;
                case EnumDirections.South:
                    return this.S;
                case EnumDirections.East:
                    return this.E;
                case EnumDirections.West:
                    return this.W;
                default:
                    return null;
            }
        }
    }
}
