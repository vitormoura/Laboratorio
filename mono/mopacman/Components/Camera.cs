using Microsoft.Xna.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    public class Camera
    {
        private Vector2 origin;
        private float width;
        private float ratio;
        
        public float Ratio
        {
            get
            {
                return (ratio < 0) ? this.ratio = (MyGame.Graphics.PreferredBackBufferWidth / this.width) : this.ratio;
            }
        }

        public Vector2 Position
        {
            get { return this.origin; }
            set { this.origin = value; }
        }

        public Camera(Vector2 origin, Int32 width)
        {
            this.ratio = -1;
            this.width = width;
            this.origin = origin;
        }

        public Point TranslateToPixels(Vector2 p)
        {
            return TranslateToPixels(p.ToPoint());
        }

        public Point TranslateToPixels(Point p)
        {
            //Sistema de coordenadas XNA tradicional (x:0, y:0 em TOP/LEFT)
            //*
            float ratio = this.Ratio;
            int x = (int)((p.X + this.Position.X) * ratio);
            int y = (int)((p.Y + this.Position.Y) * ratio);
            //*/

            //Inverter sistema de coordenadas para um plano cartesiano tradicional (x:0, y:0 em BOTTOM/LEFT)
            /*
            int x = (int)((p.X - this.Position.X) * ratio);
            int y = (int)((p.Y - this.Position.Y) * ratio);

            y = this.game.Graphics.PreferredBackBufferHeight - y;
            //*/

            return new Point(x, y);
        }

        public Rectangle TranslateToPixelsRect(Rectangle original)
        {
            float ratio = this.Ratio;
            
            int width = width = (int)(original.Width * ratio);
            int height = (int)(original.Height * ratio);
            Point pos = this.TranslateToPixels(original.Location);

            return new Rectangle(pos.X, pos.Y, width, height);
        }
    }
}
