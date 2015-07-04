using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using Microsoft.Xna.Framework.Input;
using mopacman.Animations;
using mopacman.Components;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Scenes
{
    class TitleScreenScene : GameScene
    {
        public TitleScreenScene(MyGame g)
            : base(g)
        {
        }

        public override void Initialize()
        {
            base.Initialize();
        }

        protected override void LoadContent()
        {
            this.background = this.Game.Content.Load<Texture2D>(@"Backgrounds\title_screen.png");
            this.font = this.Game.Content.Load<SpriteFont>(@"Fonts\padrao");

            this.pressStartMsg = new Text(this.Game as MyGame, this.font, "press start");
            this.pressStartBlink = new BlinkAnimation<Text>(this.pressStartMsg);

            base.LoadContent();
        }

        public override void Update(GameTime gameTime)
        {
            if (this.Enabled)
            {
                this.pressStartBlink.Update(gameTime);

                base.Update(gameTime);
            }
        }

        public override void Draw(GameTime gameTime)
        {
            if (this.Enabled && this.Visible)
            {
                MyGame.SpriteBatch.Draw(this.background, destinationRectangle: MyGame.Camera.TranslateToPixelsRect(this.background.Bounds));

                if (pressStartMsg.Visible)
                {
                    this.pressStartMsg.Draw(gameTime);
                }

                base.Draw(gameTime);
            }
        }

        private Texture2D background;
        private Text pressStartMsg;
        private BlinkAnimation<Text> pressStartBlink;
        private SpriteFont font;
    }
}
